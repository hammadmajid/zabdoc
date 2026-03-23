import type { DataStructure } from "./data/data";

interface AttendanceRecord {
	lecture: string;
	date: string;
	status: string;
}

export interface MarkEntry {
	head: string;
	max: string;
	obtained: string;
}

export type MarkCategory = "quiz" | "assignment" | "lab" | "project" | "exam" | "other";

export interface MarkGroup {
	category: MarkCategory;
	label: string;
	entries: MarkEntry[];
	totalMax: number;
	totalObtained: number | null;
}

export interface CourseData {
	courseName: string;
	instructor: string;
	records: AttendanceRecord[];
	marks: MarkEntry[];
}

// Used by form store
export interface Student {
	id: number;
	name: string;
	regNo: string;
}

export interface DocumentInfo {
	type: string;
	marks: string;
	number: string;
	date: string;
}

export interface ImageItem {
	id: string;
	file: File;
	previewUrl: string;
}

export interface FormState {
	// Student info
	studentName: string;
	regNo: string;
	isMultiMode: boolean;
	students: Student[];
	studentCounter: number;

	// Course info
	selectedClass: keyof DataStructure | "";
	selectedCourse: string;

	// Document info
	document: DocumentInfo;

	// Images
	images: FileList | null;
}
