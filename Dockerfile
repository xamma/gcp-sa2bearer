FROM gcr.io/distroless/static-debian11:nonroot

COPY gcp-sa2bearer /gcp-sa2bearer

ENTRYPOINT ["/gcp-sa2bearer"]