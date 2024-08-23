import fastify, { FastifyInstance } from "fastify";
import { env } from "../src/env";
import routes from "../src/index";

// Instantiate the fastify server
const server: FastifyInstance = fastify({ logger: true });

server.register(routes, { prefix: "/" });

// Start the server
try {
  server.listen({ port: env.PORT ?? 5000 });
} catch (err) {
  server.log.error(err);
  process.exit(1);
}
