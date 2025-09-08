const data = {
    "BsCS-4C": {
        "Database Systems": {instructor: "Dr. Shahzad Latif", code: "CSC 2203"},
        "Finite Automata Theory and Formal Languages": {instructor: "Dr. Shahzad Latif", code: "CSC 2204"},
        "Linear Algebra": {instructor: "Dr. Shahzad Latif", code: "CSC 2206"},
        "Design and Analysis of Algorithms": {instructor: "Dr. Shahzad Latif", code: "CSC 3202"},
    }
};

// Set default date to today
const dateInput = document.getElementById("date");
if (dateInput) {
    const today = new Date();
    const yyyy = today.getFullYear();
    const mm = String(today.getMonth() + 1).padStart(2, '0');
    const dd = String(today.getDate()).padStart(2, '0');
    dateInput.value = `${yyyy}-${mm}-${dd}`;
}

// Get form elements
const classSelect = document.getElementById("class");
const courseSelect = document.getElementById("course");
const instructorInput = document.getElementById("instructor");
const courseCodeInput = document.getElementById("courseCode");

// Populate class dropdown with all classes from data
Object.keys(data).forEach(function (className) {
    const option = document.createElement("option");
    option.value = className;
    option.textContent = className;
    classSelect.appendChild(option);
});

// When class is selected, populate courses for that class
classSelect.addEventListener("change", function () {
    // Clear course dropdown
    courseSelect.innerHTML = '<option value="">Select course</option>';

    // Clear instructor and course code
    instructorInput.value = "";
    courseCodeInput.value = "";

    // If a class is selected, populate courses
    if (classSelect.value && data[classSelect.value]) {
        const courses = data[classSelect.value];
        Object.keys(courses).forEach(function (courseName) {
            const option = document.createElement("option");
            option.value = courseName;
            option.textContent = courseName;
            courseSelect.appendChild(option);
        });
    }
});

// When course is selected, autofill instructor and course code
courseSelect.addEventListener("change", function () {
    // Clear instructor and course code
    instructorInput.value = "";
    courseCodeInput.value = "";

    // If both class and course are selected, fill in the details
    if (classSelect.value && courseSelect.value) {
        const courseData = data[classSelect.value][courseSelect.value];
        instructorInput.value = courseData.instructor;
        courseCodeInput.value = courseData.code;
    }
});
