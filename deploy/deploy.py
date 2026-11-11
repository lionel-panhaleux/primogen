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

REPO = Path(__file__).resolve().parent.parent
DOMAIN = "mcp.krcg.org"
NAME = "primogen"
PORT = 8040
OPT = f"/opt/{NAME}"

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

unit = files.put(
    name="Install primogen.service",
    src=StringIO(f"""[Unit]
Description=primogen, the VTES councilor's MCP server
After=network-online.target
Wants=network-online.target

[Service]
ExecStart={binary} serve -addr 127.0.0.1:{PORT} -bundle /var/lib/{NAME}/bundle.json.gz
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
    _if=lambda: unit.did_change() and not started.did_change(),
)
server.shell(
    name="Drop older builds",
    commands=[f"find {OPT} -mindepth 1 -maxdepth 1 ! -name {digest} -exec rm -rf {{}} +"],
    _if=any_changed(unit),
)
nginx_site(site=NAME, domain=DOMAIN, type="proxy", upstream=f"http://127.0.0.1:{PORT}")
