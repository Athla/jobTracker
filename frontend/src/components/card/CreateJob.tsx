import React, { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

interface JobFormData {
  name: string;
  source: string;
  description: string;
  createdat: string;
}

const CreateJob: React.FC = () => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [formData, setFormData] = useState<JobFormData>({
        name: '',
        source: '',
        description: '',
        createdat: '',
    });
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);

    const handleChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ): void => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent): Promise<void> => {
        e.preventDefault();
        setIsSubmitting(true);

        const newJob = {
            ...formData,
            createdat: new Date(formData.createdat).toISOString(),
        };

        try {
            const response = await fetch('http://localhost:8080/api/jobs/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(newJob),
            });

            if (!response.ok) {
                throw new Error('Failed to create job.');
            }

            const data = await response.json();
            console.log(data);
            
            setFormData({
                name: '',
                source: '',
                description: '',
                createdat: '',
            });
            setIsOpen(false);
        } catch (error) {
            console.error('Error', error);
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div>
            <motion.button 
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                onClick={() => setIsOpen(true)}
                className="bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 text-white font-bold py-3 px-6 rounded-lg shadow-lg transform transition-all duration-200"
            >
                Create New Job
            </motion.button>

            <AnimatePresence>
                {isOpen && (
                    <motion.div 
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        exit={{ opacity: 0 }}
                        className="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm flex items-center justify-center p-4"
                    >
                        <motion.div 
                            initial={{ scale: 0.9, opacity: 0, y: 20 }}
                            animate={{ scale: 1, opacity: 1, y: 0 }}
                            exit={{ scale: 0.9, opacity: 0, y: 20 }}
                            transition={{ type: "spring", duration: 0.5 }}
                            className="bg-white rounded-xl shadow-2xl w-full max-w-md border border-gray-100"
                        >
                            <div className="p-6 border-b border-gray-100">
                                <div className="flex justify-between items-center">
                                    <h2 className="text-2xl font-bold text-gray-800">Create New Job</h2>
                                    <motion.button 
                                        whileHover={{ scale: 1.1, rotate: 90 }}
                                        whileTap={{ scale: 0.9 }}
                                        onClick={() => setIsOpen(false)}
                                        className="text-gray-500 hover:text-gray-700 transition-colors"
                                    >
                                        <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                                        </svg>
                                    </motion.button>
                                </div>
                            </div>
                            
                            <form onSubmit={handleSubmit} className="p-6 space-y-6">
                                <div className="form-group relative">
                                    <label htmlFor="name" className="absolute -top-2.5 left-2 bg-white px-2 text-sm font-medium text-gray-600">
                                        Name
                                    </label>
                                    <input 
                                        id="name"
                                        name="name"
                                        type="text" 
                                        value={formData.name} 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-3 rounded-lg border-2 border-gray-200 focus:border-blue-500 focus:ring-0 transition-colors duration-200"
                                        placeholder="Enter job name"
                                        required 
                                    />
                                </div>

                                <div className="form-group relative">
                                    <label htmlFor="source" className="absolute -top-2.5 left-2 bg-white px-2 text-sm font-medium text-gray-600">
                                        Source
                                    </label>
                                    <input 
                                        id="source"
                                        name="source"
                                        type="text" 
                                        value={formData.source} 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-3 rounded-lg border-2 border-gray-200 focus:border-blue-500 focus:ring-0 transition-colors duration-200"
                                        placeholder="Enter source"
                                        required 
                                    />
                                </div>

                                <div className="form-group relative">
                                    <label htmlFor="description" className="absolute -top-2.5 left-2 bg-white px-2 text-sm font-medium text-gray-600">
                                        Description
                                    </label>
                                    <textarea 
                                        id="description"
                                        name="description"
                                        value={formData.description} 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-3 rounded-lg border-2 border-gray-200 focus:border-blue-500 focus:ring-0 transition-colors duration-200 min-h-[120px] resize-none"
                                        placeholder="Enter description"
                                        required 
                                    />
                                </div>

                                <div className="form-group relative">
                                    <label htmlFor="createdat" className="absolute -top-2.5 left-2 bg-white px-2 text-sm font-medium text-gray-600">
                                        Create at
                                    </label>
                                    <input 
                                        id="createdat"
                                        name="createdat"
                                        type="date" 
                                        value={formData.createdat} 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-3 rounded-lg border-2 border-gray-200 focus:border-blue-500 focus:ring-0 transition-colors duration-200"
                                        required 
                                    />
                                </div>

                                <motion.button 
                                    whileHover={{ scale: 1.01 }}
                                    whileTap={{ scale: 0.99 }}
                                    type="submit" 
                                    disabled={isSubmitting}
                                    className="w-full px-6 py-4 rounded-lg bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white font-medium shadow-md disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200"
                                >
                                    {isSubmitting ? (
                                        <span className="flex items-center justify-center space-x-2">
                                            <svg className="animate-spin h-5 w-5" viewBox="0 0 24 24">
                                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none"/>
                                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                                            </svg>
                                            <span>Creating...</span>
                                        </span>
                                    ) : (
                                        'Create Job'
                                    )}
                                </motion.button>
                            </form>
                        </motion.div>
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    );
};

export default CreateJob;
