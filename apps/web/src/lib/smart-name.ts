export function smartName(data: FormData | Record<string, unknown>): string {
	// Get form values with fallbacks
	let regNo: string;
	let courseCode: string;
	let type: string;
	let number: string;

	if (data instanceof FormData) {
		regNo = data.get("regNo")?.toString() || "";
		courseCode = data.get("courseCode")?.toString() || "";
		type = data.get("type")?.toString() || "";
		number = data.get("number")?.toString() || "";
	} else {
		// Handle JSON object
		const students = data.Students as Array<{ RegNo: string }> | undefined;
		regNo = students?.[0]?.RegNo || "";
		courseCode = (data.CourseCode as string) || "";
		type = (data.DocType as string) || "";
		number = (data.Number as string) || "";
	}

	// Build filename parts
	const parts: string[] = [];

	if (regNo) {
		parts.push(regNo);
	}

	if (courseCode) {
		parts.push(courseCode);
	}

	if (type && number) {
		// Abbreviate document types
		let typeAbbr;
		switch (type.toLowerCase()) {
			case "assignment":
				typeAbbr = "A";
				break;
			case "lab task":
				typeAbbr = "LT";
				break;
			case "lab project":
				typeAbbr = "LP";
				break;
			default:
				typeAbbr = type.charAt(0).toUpperCase();
		}
		parts.push(`${typeAbbr}${number}`);
	} else if (type) {
		// Just add type if number is not available
		parts.push(type.replace(/\s+/g, ""));
	} else if (number) {
		// Just add number if type is not available
		parts.push(`Doc${number}`);
	}

	// If we have parts, join them with spaces, otherwise use default
	return parts.length > 0 ? `${parts.join(" ")}.docx` : "document.docx";
}
