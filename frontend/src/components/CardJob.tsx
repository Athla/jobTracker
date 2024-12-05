import React from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Trash2 } from "lucide-react";

import { useState } from "react";

export interface Job {
  Id: string;
  Name: string;
  Source: string;
  Description: string;
  CreatedAt: string;
  Status?: "APPLIED" | "INTERVIEW" | "REJECTED" | "ACCEPTED";
}

interface JobCardProps {
  job: Job;
  onDelete: (id: string | number) => void;
}
const JobCard: React.FC<JobCardProps> = ({ job, onDelete }) => {
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    if (!confirm("Are you sure you want to delete this job?")) return;

    setIsDeleting(true);
    try {
      const response = await fetch(`http://localhost:8080/api/jobs/${job.Id}`, {
        method: "DELETE",
      });
      if (!response.ok) {
        throw new Error("Failed to delete job");
      }
      onDelete(job.Id);
    } catch (error) {
      console.error("Error deleting job:", error);
    } finally {
      setIsDeleting(false);
    }
  };

  const getStatusColor = (status?: string) => {
    switch (status) {
      case "APPLIED":
        return "bg-blue-100 text-blue-800";
      case "INTERVIEW":
        return "bg-yellow-100 text-yellow-800";
      case "REJECTED":
        return "bg-red-100 text-red-800";
      case "ACCEPTED":
        return "bg-green-100 text-green-800";
      default:
        return "bg-gray-100 text-gray-800";
    }
  };

  const formatDate = (dateString: string) => {
    try {
      const date = new Date(dateString);
      if (isNaN(date.getTime())) {
        return "Date not available";
      }

      return new Intl.DateTimeFormat("en-US", {
        year: "numeric",
        month: "long",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        timeZone: "UTC", // Add this line to handle UTC dates
      }).format(date);
    } catch (error) {
      console.error("Error formatting date:", error);
      return "Date not available";
    }
  };

  return (
    <Card className="relative transition-all duration-300 ease-in-out hover:scale-105 hover:shadow-lg hover:-translate-y-1 hover:bg-card/80 backdrop-blur-sm">
      <CardHeader>
        <div className="flex justify-between items-start">
          <div className="transition-transform duration-200 ease-in-out">
            <CardTitle className="transition-colors duration-200 hover:text-primary">
              {job.Name}
            </CardTitle>
            {job.Status && (
              <span
                className={`inline-block px-2 py-1 text-xs font-semibold rounded-full mt-2 transition-all duration-200 ease-in-out hover:scale-105 ${getStatusColor(
                  job.Status,
                )}`}
              >
                {job.Status}
              </span>
            )}
          </div>
          <Button
            variant="ghost"
            size="icon"
            onClick={handleDelete}
            disabled={isDeleting}
            className="h-8 w-8 text-muted-foreground transition-colors duration-200 hover:text-destructive hover:bg-destructive/10 hover:scale-110"
          >
            <Trash2 className="h-4 w-4" />
          </Button>
        </div>
        <CardDescription className="transition-opacity duration-200 group-hover:opacity-90">
          {formatDate(job.CreatedAt)}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="space-y-2">
          <p className="text-sm font-medium transition-colors duration-200 hover:text-primary/80">
            Source: {job.Source}
          </p>
          <p className="text-sm transition-colors duration-200 hover:text-foreground/90">
            {job.Description}
          </p>
        </div>
      </CardContent>
    </Card>
  );
};
export default JobCard;
