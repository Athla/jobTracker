import * as z from "zod";

export const loginSchema = z.object({
  username: z.string().min(3, "Username must be at least 3 characters"),
  password: z.string().min(6, "Password must be at least 6 characters"),
});

export const jobSchema = z.object({
  name: z.string().min(1, "Job title is required"),
  company: z.string().min(1, "Company name is required"),
  source: z.string().min(1, "Source is required"),
  description: z.string().nullable().optional(),
  job_type: z.enum([
    "FULL_TIME",
    "PART_TIME",
    "CONTRACT",
    "INTERNSHIP",
    "REMOTE",
  ]),
  status: z.enum([
    "WISHLIST",
    "APPLIED",
    "PHONE_SCREEN",
    "TECHNICAL_INTERVIEW",
    "ONSITE",
    "OFFER",
    "REJECTED",
    "ACCEPTED",
    "WITHDRAWN",
  ]),
});
