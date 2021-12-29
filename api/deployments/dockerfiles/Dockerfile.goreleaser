FROM alpine

COPY horusec-api /

# migration bin name is used in horusec operator, any change should be reflected there.
COPY horusec-api-migrate-v1-v2 /

ENTRYPOINT ["./horusec-api"]