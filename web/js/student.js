document.addEventListener("DOMContentLoaded", function () {
    const studentNameInput = document.getElementById("studentName");
    const regNoInput = document.getElementById("regNo");
    const classInput = document.getElementById("class");

    // Restore values from localStorage
    if (studentNameInput && localStorage.getItem("studentName")) {
        studentNameInput.value = localStorage.getItem("studentName");
    }
    if (regNoInput && localStorage.getItem("regNo")) {
        regNoInput.value = localStorage.getItem("regNo");
    }
    if (classInput && localStorage.getItem("class")) {
        classInput.value = localStorage.getItem("class");
        classInput.dispatchEvent(new Event('change'));
    }

    // Save changes to localStorage
    if (studentNameInput) {
        studentNameInput.addEventListener("input", function () {
            localStorage.setItem("studentName", studentNameInput.value);
        });
    }
    if (regNoInput) {
        regNoInput.addEventListener("input", function () {
            localStorage.setItem("regNo", regNoInput.value);
        });
    }
    if (classInput) {
        classInput.addEventListener("change", function () {
            localStorage.setItem("class", classInput.value);
        });
    }
});
