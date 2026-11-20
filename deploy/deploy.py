import hashlib
import os
import socket
import subprocess
import tempfile
from io import StringIO
from pathlib import Path

from pyinfra import host
from pyinfra.facts.server import Arch
from pyinfra.operations import files, server, systemd
from pyinfra.operations.util import any_changed
from server_setup import nginx_site
from server_setup.secrets import load, put_secret

REPO = Path(__file__).resolve().parent.parent
DOMAIN = "mcp.krcg.org"
NAME = "primogen"
# frankfurt has no port registry: 8040 is the engine's (zero-vtes), 8030 the forum's, 7007-9008 archon's
PORT = 8050
OPT = f"/opt/{NAME}"
# the GitHub App filing vtes_feedback's reports; its key is in secrets.sops.yaml
APP_CLIENT_ID = "Iv23liWtGTsLHMrGaGvD"

# A certificate request for a name that does not resolve here fails, and Let's Encrypt counts
# every failed validation against its hourly limit.
try:
    ip = socket.gethostbyname(DOMAIN)
except socket.gaierror:
    ip = None
if ip != host.data.ssh_hostname:
    raise RuntimeError(f"{DOMAIN} resolves to {ip}, not {host.data.ssh_hostname}")

goarch = {"x86_64": "amd64", "aarch64": "arm64"}[host.get_fact(Arch)]
dist = Path(tempfile.mkdtemp())
subprocess.run(
    ["go", "build", "-trimpath", "-o", str(dist / NAME), "."],
    cwd=REPO,
    check=True,
    env={**os.environ, "GOOS": "linux", "GOARCH": goarch, "CGO_ENABLED": "0"},
)
digest = hashlib.sha256((dist / NAME).read_bytes()).hexdigest()[:16]
# pyinfra diffs any file it replaces, and a binary does not diff: each build lands in a new directory.
binary = f"{OPT}/{digest}/{NAME}"
files.put(name="Upload primogen", src=str(dist / NAME), dest=binary, mode="755")
secrets = load(str(REPO / "secrets.sops.yaml"))
key = put_secret(name="GitHub App key", content=secrets["github_app_key"], dest=f"/etc/{NAME}/github-app.pem")
archon_args = ""
archon_cred = ""
archon_key = None
if secrets.get("archon_client_id") and secrets.get("archon_client_secret"):
    archon_key = put_secret(name="Archon OAuth secret", content=secrets["archon_client_secret"], dest=f"/etc/{NAME}/archon-secret")
    archon_args = f" -archon-client {secrets['archon_client_id']} -archon-secret ${{CREDENTIALS_DIRECTORY}}/archon-secret"
    archon_cred = f"\nLoadCredential=archon-secret:/etc/{NAME}/archon-secret"

unit = files.put(
    name="Install primogen.service",
    src=StringIO(f"""[Unit]
Description=primogen, the VTES councilor's MCP server
After=network-online.target
Wants=network-online.target

[Service]
ExecStart={binary} serve -addr 127.0.0.1:{PORT} -bundle /var/lib/{NAME}/bundle.json.gz -github-app {APP_CLIENT_ID} -github-key ${{CREDENTIALS_DIRECTORY}}/github-app.pem{archon_args}
LoadCredential=github-app.pem:/etc/{NAME}/github-app.pem{archon_cred}
DynamicUser=yes
StateDirectory={NAME}
Restart=always
RestartSec=5
SyslogIdentifier={NAME}

[Install]
WantedBy=multi-user.target
"""),
    dest=f"/etc/systemd/system/{NAME}.service",
    mode="644",
)
systemd.daemon_reload(name="Reload units for primogen", _if=any_changed(unit))
started = systemd.service(name="primogen running", service=f"{NAME}.service", running=True, enabled=True)
# a first start already runs the new unit, and a restart would cut its first build short
systemd.service(
    name="Restart primogen",
    service=f"{NAME}.service",
    restarted=True,
    _if=lambda: (unit.did_change() or key.did_change() or (archon_key is not None and archon_key.did_change())) and not started.did_change(),
)
server.shell(
    name="Drop older builds",
    commands=[f"find {OPT} -mindepth 1 -maxdepth 1 ! -name {digest} -exec rm -rf {{}} +"],
    _if=any_changed(unit),
)
nginx_site(site=NAME, domain=DOMAIN, type="proxy", upstream=f"http://127.0.0.1:{PORT}")
