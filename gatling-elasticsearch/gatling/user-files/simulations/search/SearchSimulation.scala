package search

import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._

class SearchSimulation extends Simulation {

  val baseUrl = "http://localhost:9200"

  val httpProtocol = http
    // Here is the root for all relative URLs
    .baseUrl(baseUrl)
    // Here are the common headers
    .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    .doNotTrackHeader("1")
    .acceptLanguageHeader("en-US,en;q=0.5")
    .acceptEncodingHeader("gzip, deflate")
    .userAgentHeader("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:16.0) Gecko/20100101 Firefox/16.0")

  var searchFeeder = csv("search.csv").circular

  // A scenario is a chain of requests and pauses
  val scn = scenario("Scenario Name")
    .feed(searchFeeder)
    .exec(
      http("request_1")
        .get("/shakespeare/_search")
        .header("Content-Type", "application/json")
        .body(
          StringBody(
            """{"query":{"match":{"text_entry":"${keyword}"}}}"""
          )
        ).asJson
    )

  setUp(scn.inject(constantUsersPerSec(3) during (10 seconds)).protocols(httpProtocol))
}
