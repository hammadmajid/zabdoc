export function smartName(formData: FormData): string {
	// Get form values with fallbacks
	const regNo = formData.get("regNo")?.toString() || "";
	const courseCode = formData.get("courseCode")?.toString() || "";
	const type = formData.get("type")?.toString() || "";
	const number = formData.get("number")?.toString() || "";

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
	return parts.length > 0 ? `${parts.join(" ")}.pdf` : "document.pdf";
}
