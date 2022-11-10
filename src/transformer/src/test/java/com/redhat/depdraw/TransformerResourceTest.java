package com.redhat.depdraw;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

@QuarkusTest
public class TransformerResourceTest {

    @Test
    public void testTransformerHealthEndpoint() {
        given()
          .when().get("/health")
          .then()
             .statusCode(200);
    }

}