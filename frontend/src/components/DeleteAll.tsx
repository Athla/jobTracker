import React from "react";
import { Button } from "./ui/button";

interface DeleteAllProps {
  onDeleteAll: () => void;
  token: string;
}

const DeleteAll: React.FC<DeleteAllProps> = ({ onDeleteAll, token }) => {
  const handleDeleteAll = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/jobs/deleteAll", {
        method: "DELETE",
        headers: {
          Authorization: `${token}`,
        },
      });
      if (!response.ok) {
        throw new Error("Failed to delete all jobs");
      }
      onDeleteAll(); // Call the parent function to update the state
    } catch (error) {
      console.error("Error deleting all jobs", error);
    }
  };

  return (
    <Button
      variant="destructive"
      onClick={handleDeleteAll}
      className="hover:bg-destructive/90"
    >
      Delete All Jobs
    </Button>
  );
};

export default DeleteAll;
