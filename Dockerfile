FROM scratch

ADD ./cassandra ./cassandra

CMD ["./cassandra"]
