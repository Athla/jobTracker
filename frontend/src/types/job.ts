export type JobType =
  | "FULL_TIME"
  | "PART_TIME"
  | "CONTRACT"
  | "INTERNSHIP"
  | "REMOTE";

export type JobStatus =
  | "WISHLIST"
  | "APPLIED"
  | "PHONE_SCREEN"
  | "TECHNICAL_INTERVIEW"
  | "ONSITE"
  | "OFFER"
  | "REJECTED"
  | "ACCEPTED"
  | "WITHDRAWN"
  | "INTERVIEWING";

export interface Job {
  id: string;
  name: string;
  company: string;
  source: string;
  description?: string | null;
  job_type: JobType;
  status: JobStatus;
  version: number;
  created_at: string;
  updated_at: string;
}
