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
data class Verification(val type: String, val token: String, val challenge: String = "")


fun Application.configureSerialization() {
    install(ContentNegotiation) {
        json(
            Json {
                ignoreUnknownKeys = true
            }
        )
    }
    
    val categories = mapOf("fruits" to arrayOf("banana", "apple", "orange"), 
                           "vegetables" to arrayOf("salad", "onion", "potato", "broccoli"))

    routing {
        post("/slack/events") {
                val json = call.receive<Verification>()
                call.respond(json.challenge)
        }

        post("/slack/ping") {
                call.respondText("pong")
        }
        
        post("/slack/categories") {
                call.respondText(categories.keys.joinToString())
        }
        
        post("/slack/categories/list") {
                val formParameters = call.receiveParameters()
                println(formParameters)
                val category = formParameters["text"].toString()
                if(! categories.containsKey(category)) {
                    call.respondText("Incorrect category: $category")    
                }
                call.respondText(categories.getValue(category).joinToString())
        }
    }
}
