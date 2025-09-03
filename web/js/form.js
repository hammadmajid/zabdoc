const data = {
    "BsCS-4C": {
        "Automota Theory": "Dr. Shezad Latif",
        "Database Management Systems": "Qanetah Ahmed"
    }
};

const classSelect = document.getElementById("class");
const courseSelect = document.getElementById("course");
const instructorSelect = document.getElementById("instructor");

function safeGetElement(el, name) {
    if (!el) {
        throw new Error(`Element with id "${name}" not found in DOM`);
    }
    return el;
}

function safeGetClass(cls) {
    if (!(cls in data)) {
        throw new Error(`Invalid class selected: "${cls}"`);
    }
    return data[cls];
}

function safeGetCourse(cls, course) {
    const courses = safeGetClass(cls);
    if (!(course in courses)) {
        throw new Error(`Invalid course "${course}" for class "${cls}"`);
    }
    return courses[course];
}

(function init() {
    try {
        // Set default date to today
        const dateInput = document.getElementById("date");
        if (dateInput) {
            const today = new Date();
            const yyyy = today.getFullYear();
            const mm = String(today.getMonth() + 1).padStart(2, '0');
            const dd = String(today.getDate()).padStart(2, '0');
            dateInput.value = `${yyyy}-${mm}-${dd}`;
        }

        const cSelect = safeGetElement(classSelect, "class");
        const crsSelect = safeGetElement(courseSelect, "course");
        const instSelect = safeGetElement(instructorSelect, "instructor");

        Object.keys(data).forEach(cls => {
            let option = document.createElement("option");
            option.value = cls;
            option.textContent = cls;
            cSelect.appendChild(option);
        });

        cSelect.addEventListener("change", () => {
            try {
                crsSelect.innerHTML = '<option value="">Select course</option>';
                instSelect.innerHTML = '<option value="">Select instructor</option>';

                if (cSelect.value) {
                    const courses = safeGetClass(cSelect.value);
                    Object.keys(courses).forEach(course => {
                        let option = document.createElement("option");
                        option.value = course;
                        option.textContent = course;
                        crsSelect.appendChild(option);
                    });
                }
            } catch (err) {
                console.error(err.message);
            }
        });

        crsSelect.addEventListener("change", () => {
            try {
                instSelect.value = "";
                if (cSelect.value && crsSelect.value) {
                    instSelect.value = safeGetCourse(cSelect.value, crsSelect.value);
                }
            } catch (err) {
                console.error(err.message);
            }
        });

    } catch (err) {
        console.error(err.message);
    }
})();
