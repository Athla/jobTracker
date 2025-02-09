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
  | "WITHDRAWN";

export interface Job {
  id: string;
  name: string;
  source: string;
  description: string;
  company: string;
  location?: string;
  salary_range?: string;
  job_type: JobType;
  status: JobStatus;
  application_link?: string;
  rejection_reason?: string;
  notes?: string;
  interview_notes?: string;
  next_interview_date?: string;
  last_interaction_date?: string;
  version: number;
  created_at: string;
  updated_at: string;
}
