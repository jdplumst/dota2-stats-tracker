import z from "zod";

const envSchema = z.object({
  NODE_ENV: z
    .enum(["development", "production", "test"])
    .default("development"),
  PORT: z.number().default(5000),
});

export const env = envSchema.parse(process.env);
