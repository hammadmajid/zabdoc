import json
import re
from bs4 import BeautifulSoup

def parse_html_to_ts(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        html_content = f.read()

    soup = BeautifulSoup(html_content, 'html.parser')
    data = {}

    tables = soup.find_all('table', align="center", border="1")
    
    for table in tables:
        rows = table.find_all('tr')
        if len(rows) < 2:
            continue

        heading_td = rows[0].find('td', class_='tableHeading')
        if not heading_td:
            continue
            
        heading_text = heading_td.get_text(strip=True)
        
        section_key = ""
        match_section = re.search(r"BCS/BS\s+(\d)\s+-\s*Section\s+([A-Z])", heading_text)
        match_open = re.search(r"-\s*Open", heading_text)

        if match_section:
            section_key = f"BsCS-{match_section.group(1)}{match_section.group(2)}"
        elif match_open:
            section_key = "Open"
        else:
            continue

        if section_key not in data:
            data[section_key] = {}

        for i in range(2, len(rows)):
            cols = rows[i].find_all('td')
            if len(cols) >= 3:
                raw_course = cols[1].get_text(strip=True)
                
                if "final year project" in raw_course.lower():
                    continue

                raw_faculty = cols[2].get_text(strip=True)

                raw_faculty = raw_faculty.replace('\xa0', '')
                raw_faculty = re.sub(r'\s+', ' ', raw_faculty)
                raw_faculty = re.sub(r'\s+[-.]\s*$', '', raw_faculty).strip()

                course_match = re.match(r"^([A-Z]+\s+\d+)\s+(.+?)(?:\s*\(\d,\d\))?$", raw_course)
                if course_match:
                    code = course_match.group(1).strip()
                    name = course_match.group(2).strip()
                    
                    data[section_key][name] = {
                        "instructor": raw_faculty,
                        "code": code
                    }

    ts_output = (
        "export type CourseInfo = {\n"
        "    instructor: string;\n"
        "    code: string;\n"
        "};\n\n"
        "export type ClassData = Record<string, CourseInfo>;\n\n"
        "export type DataStructure = Record<string, ClassData>;\n\n"
        f"export const data: DataStructure = {json.dumps(data, indent=4)};\n"
    )

    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(ts_output)

if __name__ == "__main__":
    parse_html_to_ts('input.html', '../worker/src/lib/data/data.ts')
