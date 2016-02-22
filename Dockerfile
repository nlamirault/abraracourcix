# AUTHOR:         Nicolas Lamirault <nicolas.lamirault@gmail.com>
# DESCRIPTION:    abraracourcix Dockerfile

FROM scratch
MAINTAINER Nicolas Lamirault <nicolas.lamirault@gmail.com>

COPY abraracourcix /
EXPOSE 80
ENTRYPOINT ["/abraracourcix"]
