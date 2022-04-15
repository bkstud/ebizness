package com.example.plugins

import io.ktor.serialization.kotlinx.json.*
import io.ktor.server.plugins.*
import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.request.*
import io.ktor.server.routing.*
import io.ktor.server.plugins.contentnegotiation.*
import kotlinx.serialization.*
import kotlinx.serialization.json.*


@Serializable
data class Verification(val type: String, val token: String, val challenge: String)
// @Serializable
// data class Body(val body: Verification)

fun Application.configureSerialization() {
    install(ContentNegotiation) {
        json(
            Json {
                ignoreUnknownKeys = true
            }
        )
    }

    routing {
        post("/slack/events") {
                val json = call.receive<Verification>()
                println(json)
                call.respond(json.challenge)
            }
    }
}
