import { browser } from "$app/environment";
import { data, type DataStructure, type CourseInfo } from "$lib/data/data";
import type { DocumentInfo, Student } from "$lib/types";
import { wizardStore } from "./wizard-store.svelte";
import { toast } from "svelte-sonner";
import { documentSchema, type DocumentSchema } from "@repo/types";

// Constants
export const STUDENT_LIMIT = 6;

// LocalStorage utility functions
function setLocalStorage(key: string, value: string) {
	if (!browser) return;
	try {
		localStorage.setItem(key, value);
	} catch (error) {
		toast.warning("Failed to save to localStorage", {
			description: error as string,
			position: "top-center"
		});
	}
}

function getLocalStorage(key: string): string {
	if (!browser) return "";
	try {
		return localStorage.getItem(key) || "";
	} catch (error) {
		toast.warning("Failed to read from localStorage", {
			description: error as string,
			position: "top-center"
		});
		return "";
	}
}

// Create the form store
function createFormStore() {
	// Initialize state with defaults
	let studentName = $state("");
	let regNo = $state("");
	let isMultiMode = $state(false);
	let isBlankMode = $state(false);
	let students = $state<Student[]>([
		{ id: 1, name: "", regNo: "" },
		{ id: 2, name: "", regNo: "" }
	]);
	let studentCounter = $state(2);

	let selectedClass = $state<keyof DataStructure | "">("");
	let selectedCourse = $state("");

	let document = $state<DocumentInfo>({
		type: "",
		marks: "",
		number: "",
		date: ""
	});

	// Derived values for course selection
	const classes = $derived(
		Object.keys(data).map((className) => ({
			value: className,
			label: className
		}))
	);

	const courses = $derived(
		selectedClass && selectedClass in data
			? Object.keys(data[selectedClass as keyof DataStructure])
					.filter((courseName) => {
						if (wizardStore.documentType === "Assignment") {
							return !courseName.startsWith("Lab:");
						} else if (wizardStore.documentType === "Lab Task") {
							return courseName.startsWith("Lab:");
						}
						return true; // Include all if no type selected
					})
					.map((courseName) => ({
						value: courseName,
						label: courseName
					}))
			: []
	);

	const courseDetails = $derived<CourseInfo | null>(
		selectedClass &&
			selectedCourse &&
			selectedClass in data &&
			selectedCourse in data[selectedClass]
			? data[selectedClass][selectedCourse]
			: null
	);

	const classTriggerContent = $derived(
		classes.find((c) => c.value === selectedClass)?.label ?? "Select class"
	);

	const courseTriggerContent = $derived(
		courses.find((c) => c.value === selectedCourse)?.label ?? "Select course"
	);

	// Initialize from localStorage (studentName, regNo, and selectedClass)
	function initFromLocalStorage() {
		if (browser) {
			studentName = getLocalStorage("studentName") || "";
			regNo = getLocalStorage("regNo") || "";
			const savedClass = getLocalStorage("selectedClass");
			if (savedClass && savedClass in data) {
				selectedClass = savedClass as keyof DataStructure;
			}
		}
	}

	// Persist to localStorage (only studentName and regNo)
	function persistStudentName(value: string) {
		studentName = value;
		if (browser && value) {
			setLocalStorage("studentName", value);
		}
	}

	function persistRegNo(value: string) {
		regNo = value;
		if (browser && value) {
			setLocalStorage("regNo", value);
		}
	}

	// Student management functions
	function addStudent() {
		if (students.length < STUDENT_LIMIT) {
			studentCounter++;
			students = [...students, { id: studentCounter, name: "", regNo: "" }];
		}
	}

	function removeStudent(studentId: number) {
		if (students.length > 2) {
			students = students.filter((student) => student.id !== studentId);
		}
	}

	function updateStudent(studentId: number, field: "name" | "regNo", value: string) {
		students = students.map((student) =>
			student.id === studentId ? { ...student, [field]: value } : student
		);
	}

	// Course selection functions
	function setSelectedClass(value: keyof DataStructure | "") {
		selectedClass = value;
		// Reset course when class changes
		selectedCourse = "";
		// Persist class to localStorage
		if (browser && value) {
			setLocalStorage("selectedClass", value);
		}
	}

	function setSelectedCourse(value: string) {
		selectedCourse = value;
	}

	// Validate and reset course if it's not in the current courses list
	function validateSelectedCourse() {
		if (selectedCourse && !courses.some((course) => course.value === selectedCourse)) {
			selectedCourse = "";
		}
	}

	// Document functions
	function setDocumentType(type: string) {
		document.type = type;
	}

	function setDocumentMarks(marks: string) {
		document.marks = marks;
	}

	function setDocumentNumber(num: string) {
		document.number = num;
	}

	function setDocumentDate(date: string) {
		document.date = date;
	}

	// Build JSON payload for submission matching zod schema
	function buildJSON(): DocumentSchema {
		// Build students array based on mode
		let studentsArray: { Name: string; RegNo: string }[] = [];

		if (isBlankMode) {
			// For blank mode, send empty student
			studentsArray = [{ Name: "", RegNo: "" }];
		} else if (isMultiMode) {
			// For group mode, send all students
			studentsArray = students.map((s) => ({
				Name: s.name,
				RegNo: s.regNo
			}));
		} else {
			// For individual mode, send single student
			studentsArray = [{ Name: studentName, RegNo: regNo }];
		}

		const payload = {
			Students: studentsArray,
			Class: selectedClass,
			Course: selectedCourse,
			CourseCode: courseDetails?.code || "",
			Instructor: courseDetails?.instructor || "",
			DocType: document.type,
			Number: document.number,
			Date: document.date,
			Marks: document.marks
		};

		// Validate with zod schema
		const validated = documentSchema.parse(payload);
		return validated;
	}

	// Reset form (except persisted values)
	function reset() {
		isMultiMode = false;
		isBlankMode = false;
		students = [
			{ id: 1, name: "", regNo: "" },
			{ id: 2, name: "", regNo: "" }
		];
		studentCounter = 2;
		selectedClass = "";
		selectedCourse = "";
		document = {
			type: "",
			marks: "",
			number: "",
			date: ""
		};
	}

	return {
		// Student state
		get studentName() {
			return studentName;
		},
		set studentName(value: string) {
			persistStudentName(value);
		},
		get regNo() {
			return regNo;
		},
		set regNo(value: string) {
			persistRegNo(value);
		},
		get isMultiMode() {
			return isMultiMode;
		},
		set isMultiMode(value: boolean) {
			isMultiMode = value;
		},
		get isBlankMode() {
			return isBlankMode;
		},
		set isBlankMode(value: boolean) {
			isBlankMode = value;
		},
		get students() {
			return students;
		},

		// Course state
		get selectedClass() {
			return selectedClass;
		},
		set selectedClass(value: keyof DataStructure | "") {
			setSelectedClass(value);
		},
		get selectedCourse() {
			return selectedCourse;
		},
		set selectedCourse(value: string) {
			setSelectedCourse(value);
		},

		// Document state
		get document() {
			return document;
		},

		// Derived values
		get classes() {
			return classes;
		},
		get courses() {
			return courses;
		},
		get courseDetails() {
			return courseDetails;
		},
		get classTriggerContent() {
			return classTriggerContent;
		},
		get courseTriggerContent() {
			return courseTriggerContent;
		},

		// Methods
		initFromLocalStorage,
		addStudent,
		removeStudent,
		updateStudent,
		setDocumentType,
		setDocumentMarks,
		setDocumentNumber,
		setDocumentDate,
		validateSelectedCourse,
		buildJSON,
		reset,
		STUDENT_LIMIT
	};
}

// Export singleton instance
export const formStore = createFormStore();
