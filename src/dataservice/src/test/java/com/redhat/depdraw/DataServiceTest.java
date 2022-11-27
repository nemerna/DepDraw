package com.redhat.depdraw;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;

@QuarkusTest
public class DataServiceTest {

    @Test
    public void testDataServiceHealthEndpoint() {
        given()
                .when().get("/health")
                .then()
                .statusCode(200);
    }

}