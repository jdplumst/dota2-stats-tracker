"use strict";

import fastify, { FastifyInstance } from "fastify";
import { env } from "../src/env.js";
import routes from "../src/index.js";

// Instantiate the fastify server
const server: FastifyInstance = fastify({ logger: true });

server.register(routes, { prefix: "/" });

export default async (req, res) => {
  await server.ready();
  server.server.emit("request", req, res);
};

// // Start the server
// try {
//   server.listen({ port: env.PORT ?? 5000 });
// } catch (err) {
//   server.log.error(err);
//   process.exit(1);
// }
