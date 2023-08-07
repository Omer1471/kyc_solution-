FROM alpinelinux/ansible:latest
WORKDIR /ansible
COPY . /ansible
CMD ["tail", "-f", "/dev/null"]

