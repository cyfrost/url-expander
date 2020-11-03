FROM alpine
COPY url-expander /usr/bin
ENV HOST localhost
ENV PORT 5000
EXPOSE 5000
CMD /usr/bin/url-expander
