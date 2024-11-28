# See https://stackoverflow.com/questions/76197091/quarkus-equivalent-of-wildfly-jboss-startup-scripts-in-keycloak

FROM keycloak/keycloak:26.0

WORKDIR /opt/keycloak
RUN /opt/keycloak/bin/kc.sh build

COPY ./entrypoint.sh /opt/keycloak/bin/entrypoint
ENTRYPOINT ["/opt/keycloak/bin/entrypoint"]
CMD ["start"]
