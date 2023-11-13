# Debian for playwright
FROM python:3.10.13-bookworm
LABEL maintainer="shur3d"
LABEL description="Dockerized urlwatch including playwright for browser-jobs in urlwatch"

# Do not build rust
ARG CRYPTOGRAPHY_DONT_BUILD_RUST=1

# Install all packages
RUN apt-get -y update
RUN apt-get -y upgrade
#RUN apt-get -y install python3-venv python3-pip
RUN apt-get -y install build-essential libssl-dev libpoppler-cpp-dev pkg-config python3-dev gcc musl-dev python3-dev libffi-dev cargo cron pkg-config

# Pip-install all Python modules
RUN python3 -m pip install cryptography --break-system-packages
RUN python3 -m pip install pyyaml minidb requests keyring appdirs lxml cssselect --break-system-packages
RUN python3 -m pip install chump pushbullet.py matrix_client markdown2 colorama pytest-playwright beautifulsoup4 jsbeautifier cssbeautifier pdftotext pytesseract Pillow aioxmpp jq --break-system-packages
RUN apt-get -y install libatk1.0-0 libatk-bridge2.0-0 libcups2 libdrm2 libdbus-1-3 libxkbcommon0 libatspi2.0-0 libxcomposite1 libxdamage1 libxfixes3 libxrandr2 libgbm1 libasound2

# Install initializations, applicatios and default cron file
RUN playwright install
RUN python3 -m pip install urlwatch
RUN apt-get -y install vim
RUN echo '*/30 * * * * cd /root/.urlwatch && urlwatch --urls urls.yaml --config urlwatch.yaml --hooks hooks.py --cache cache.db' | crontab -

# Cleanup
RUN rm -rf /var/lib/apt/lists/*

# Playwright configuration to use the installed chromium
#ENV CHROME_BIN=/usr/bin/chromium-browser \
#    CHROME_PATH=/usr/lib/chromium/ \
#    PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1

VOLUME /root/.urlwatch
WORKDIR /root/.urlwatch

# Keeping commandline alive and publishing tail
CMD ["tail", "-f", "/dev/stdout"]