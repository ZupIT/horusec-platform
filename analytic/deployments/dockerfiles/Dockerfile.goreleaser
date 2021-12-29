FROM alpine

COPY horusec-analytic /

# migration bin name is used in horusec operator, any change should be reflected there.
COPY horusec-analytic-migrate-v1-v2 /

ENTRYPOINT ["./horusec-analytic"]