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