## Compliment Generator Project:

---

# Compliment Generator 

A simple Go program that delivers random compliments to brighten your day. Press **ENTER** to receive a compliment, or type **quit** to exit.  

---

## Features
- **Random compliments** drawn from a curated list.  
- **Colorful output** using ANSI escape codes for a lively terminal experience.  
- **Quit option** — type `quit` to gracefully exit.  
- **Category support** (motivational, confidence, kindness) for organized compliments.  

---

## Project Structure
- `main.go` → Entry point of the program.  
- `compliments` → Slice or map containing compliment strings.  
- `Generator()` → Function to select a random compliment.  
- `UserInput()` → Handles terminal input.  

---

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/michaelbag8/compliment-generator.git
   cd compliment-generator
   ```
2. Run the program:
   ```bash
   go run main.go
   ```
3. Press **ENTER** to receive a compliment.  
4. Type **quit** to exit the program.  

---

## Example Output
```
<><> Welcome to the Compliment Generator <><>
=== Press ENTER to receive a compliment ===
Type "quit" to exit

You are doing great
```

With colors enabled:
```
<><> Welcome to the Compliment Generator <><>
=== Press ENTER to receive a compliment ===
Type "quit" to exit

Your kindness inspires others
```
*(Displayed in green text on your terminal)*

---

## Future Improvements
- Allow users to **choose categories** (e.g., motivational, confidence, kindness).  
- Add **background colors** or combined styles for compliments.  
- Support **custom compliments** entered by the user.  

---