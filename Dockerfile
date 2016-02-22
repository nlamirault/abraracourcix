# AUTHOR:         Nicolas Lamirault <nicolas.lamirault@gmail.com>
# DESCRIPTION:    abraracourcix Dockerfile

FROM scratch
MAINTAINER Nicolas Lamirault <nicolas.lamirault@gmail.com>

COPY abraracourcix /
EXPOSE 8080
ENTRYPOINT ["/abraracourcix"]
