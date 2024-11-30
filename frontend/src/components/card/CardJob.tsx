import React from 'react';

export interface Job {
  Id: string;
  Name: string;
  Source: string;
  Description: string;
  CreatedAt: string;
}

interface JobCardProps {
  job: Job;
  onDelete: (id: string | number) => void;
}

const deleteJob = async (id: string | number) => {
  try {
    const response = await fetch(`http://localhost:8080/api/jobs/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error('Failed to delete job');
    }
  } catch (error) {
    console.error('Error deleting job', error);
    throw error;
  }
};

const JobCard: React.FC<JobCardProps> = ({ job, onDelete }) => {
  const handleDelete = async () => {
    try {
      await deleteJob(job.Id);
      onDelete(job.Id);
    } catch (error) {
      console.error(`Failed to delete job due: ${error}`, error);
    }
  };

  return (
    <div className="bg-white rounded-xl shadow-lg overflow-hidden w-72 hover:shadow-xl transition-shadow duration-300">
      <div className="p-5">
        <div className="border-l-4 border-blue-500 pl-4">
          <h2 className="text-xl font-bold text-gray-800 truncate">{job.Name}</h2>
          <span className="inline-block bg-blue-100 text-blue-600 text-xs px-2 py-1 rounded-full mt-2">
            {job.Source}
          </span>
        </div>
        
        <div className="mt-4">
          <p className="text-gray-600 text-sm line-clamp-3">
            {job.Description}
          </p>
        </div>

        <div className="mt-4 flex items-center justify-between">
          <span className="text-xs text-gray-500">
            {new Date(job.CreatedAt).toLocaleDateString()}
          </span>
          <button 
            onClick={handleDelete}
            className="bg-red-50 text-red-600 px-3 py-1 rounded-lg text-sm font-medium
                     hover:bg-red-100 transition-colors duration-200"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
};

export default JobCard;
