import { browser } from "$app/environment";
import { data, type DataStructure, type CourseInfo } from "$lib/data/data";
import { toast } from "svelte-sonner";

// Types
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
            position: "top-center",
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
            position: "top-center",
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
    let students = $state<Student[]>([
        { id: 1, name: "", regNo: "" },
        { id: 2, name: "", regNo: "" },
    ]);
    let studentCounter = $state(2);

    let selectedClass = $state<keyof DataStructure | "">("");
    let selectedCourse = $state("");

    let document = $state<DocumentInfo>({
        type: "",
        marks: "",
        number: "",
        date: "",
    });

    let images = $state<FileList | null>(null);

    // Derived values for course selection
    const classes = $derived(
        Object.keys(data).map((className) => ({
            value: className,
            label: className,
        }))
    );

    const courses = $derived(
        selectedClass && selectedClass in data
            ? Object.keys(data[selectedClass]).map((courseName) => ({
                  value: courseName,
                  label: courseName,
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

    const hasImages = $derived(images && images.length > 0);

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

    // Image functions
    function setImages(fileList: FileList | null) {
        images = fileList;
    }

    // Build FormData for submission
    function buildFormData(): FormData {
        const formData = new FormData();

        // Student info
        formData.append("isMultiMode", isMultiMode.toString());

        if (isMultiMode) {
            students.forEach((student, index) => {
                formData.append(`student-${index + 1}-name`, student.name);
                formData.append(`student-${index + 1}-regNo`, student.regNo);
            });
        } else {
            formData.append("studentName", studentName);
            formData.append("regNo", regNo);
        }

        // Course info
        formData.append("class", selectedClass);
        formData.append("course", selectedCourse);

        if (courseDetails) {
            formData.append("instructor", courseDetails.instructor);
            formData.append("courseCode", courseDetails.code);
        }

        // Document info
        if (document.type) formData.append("type", document.type);
        if (document.marks) formData.append("marks", document.marks);
        if (document.number) formData.append("number", document.number);
        if (document.date) formData.append("date", document.date);

        // Images
        if (images) {
            for (let i = 0; i < images.length; i++) {
                formData.append("images", images[i]);
            }
        }

        return formData;
    }

    // Reset form (except persisted values)
    function reset() {
        isMultiMode = false;
        students = [
            { id: 1, name: "", regNo: "" },
            { id: 2, name: "", regNo: "" },
        ];
        studentCounter = 2;
        selectedClass = "";
        selectedCourse = "";
        document = {
            type: "",
            marks: "",
            number: "",
            date: "",
        };
        images = null;
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

        // Images state
        get images() {
            return images;
        },
        set images(value: FileList | null) {
            setImages(value);
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
        get hasImages() {
            return hasImages;
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
        buildFormData,
        reset,
        STUDENT_LIMIT,
    };
}

// Export singleton instance
export const formStore = createFormStore();
