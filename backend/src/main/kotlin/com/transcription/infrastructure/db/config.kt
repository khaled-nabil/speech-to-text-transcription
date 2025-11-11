package com.transcription.infrastructure.db

import io.ktor.server.application.*
import org.jetbrains.exposed.sql.*

fun Application.configureDatabases() {
    val dbHost = System.getenv("DB_HOST")
        ?: throw IllegalStateException("DB_HOST environment variable is not set")
    val dbPort = System.getenv("DB_PORT")
        ?: throw IllegalStateException("DB_PORT environment variable is not set")
    val dbName = System.getenv("DB_NAME")
        ?: throw IllegalStateException("DB_NAME environment variable is not set")
    val user = System.getenv("DB_USER")
        ?: throw IllegalStateException("DB_USER environment variable is not set")
    val password = System.getenv("DB_PASSWORD")
        ?: throw IllegalStateException("DB_PASSWORD environment variable is not set")

    val url = "jdbc:postgresql://$dbHost:$dbPort/$dbName"

    log.info("Connecting to Postgres database")

    Database.connect(
        url,
        user = user,
        password = password
    )
}