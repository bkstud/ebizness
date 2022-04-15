package com.example.plugins

import io.ktor.server.routing.*
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.request.*
import com.slack.api.*

fun Application.configureRouting() {

    routing {
        get("/") {
            val token = System.getenv("SLACK_TOKEN")
            val slack = Slack.getInstance()
            val response = slack.methods(token).chatPostMessage {
                it.channel("#general")
                .text("Hello :wave:")
            }
            call.respondText("Response is: $response")
        }
        
    }
}
