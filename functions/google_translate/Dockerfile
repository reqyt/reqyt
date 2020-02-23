FROM alpine
RUN apk add libc6-compat
ADD ./google_translate /
HEALTHCHECK --interval=1s --timeout=3s CMD wget -O - -q localhost:8000
CMD ["/google_translate"]
EXPOSE 8000
