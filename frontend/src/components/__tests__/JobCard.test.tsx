import { render, screen } from "@/test/utils";
import JobCard from "../JobCard";
import { Job } from "@/types/job";

const mockJob: Job = {
  id: "1",
  name: "Software Engineer",
  company: "Tech Corp",
  source: "LinkedIn",
  description: "Great job opportunity",
  job_type: "FULL_TIME",
  status: "APPLIED",
  version: 1,
  created_at: "2024-01-01T00:00:00Z",
  updated_at: "2024-01-01T00:00:00Z",
};

describe("JobCard", () => {
  it("renders job information correctly", () => {
    render(<JobCard job={mockJob} />);

    // Check if main job details are rendered
    expect(screen.getByText("Software Engineer")).toBeInTheDocument();
    expect(screen.getByText("Tech Corp")).toBeInTheDocument();
    expect(screen.getByText("LinkedIn")).toBeInTheDocument();
    expect(screen.getByText("Great job opportunity")).toBeInTheDocument();

    // Check if status badge is rendered
    expect(screen.getByText("APPLIED")).toBeInTheDocument();

    // Check if date is formatted correctly
    expect(screen.getByText(/Jan 1, 2024/)).toBeInTheDocument();
  });

  it("applies correct status color classes", () => {
    render(<JobCard job={mockJob} />);

    const statusBadge = screen.getByText("APPLIED");
    expect(statusBadge).toHaveClass("bg-yellow-100", "text-yellow-800");
  });
});
