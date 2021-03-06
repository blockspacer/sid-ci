FROM python:3.8

RUN groupadd --gid 2000 sid-daemon && useradd --create-home sid-daemon --gid 2000 --uid 1000

ARG YOUR_ENV

ENV YOUR_ENV=${YOUR_ENV} \
  PYTHONFAULTHANDLER=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONHASHSEED=random \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PIP_DEFAULT_TIMEOUT=100 \
  POETRY_VERSION=1.0.0

# System deps:
RUN pip install "poetry==$POETRY_VERSION"

# Copy only requirements to cache them in docker layer
WORKDIR /code
COPY poetry.lock pyproject.toml /code/

# Project initialization:
RUN poetry config virtualenvs.create false \
  && poetry install $(test "$YOUR_ENV" == production && echo "--no-dev") --no-interaction --no-ansi

RUN git config --global user.email "sid@sidci.io" && git config --global user.name "Dev McCode"

# Creating folders, and files for a project:
COPY . /code
RUN chmod +x git-exe.sh
ENV GIT_SSH_VARIANT=ssh


CMD [ "python", "main.py" ]