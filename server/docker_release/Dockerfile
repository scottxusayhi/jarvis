FROM dev.k2data.com.cn:5001/alpine
COPY jarvis_master_linux_amd64.bin /opt/k2data/jarvis/jarvis_master
COPY web/build /opt/k2data/jarvis/web/build
WORKDIR /opt/k2data/jarvis
EXPOSE 2999

ARG branch
ARG commit
ARG buildtime
ARG owner
ARG env_para
LABEL branch=$branch \
        commit=$commit \
        buildtime=$buildtime \
        owner=$owner \
        env_para=$env_para
# put your commands here

ENTRYPOINT ["./jarvis_master"]
CMD ["-h"]
