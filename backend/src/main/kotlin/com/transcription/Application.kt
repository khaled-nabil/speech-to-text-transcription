package com.transcription

import com.transcription.infrastructure.db.configureDatabases
import com.transcription.presentation.routes.configureRouting
import io.ktor.server.application.*

fun main(args: Array<String>) {
    io.ktor.server.netty.EngineMain.main(args)
}

fun Application.module() {
    configureDatabases()
    configureRouting()
}
