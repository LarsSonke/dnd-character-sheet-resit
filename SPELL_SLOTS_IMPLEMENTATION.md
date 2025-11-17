# Spell Slot Usage Tracking Implementation

## Summary
Implemented spell slot usage tracking for D&D 5e character sheets, allowing characters to cast spells which consume the appropriate spell slots. Cantrips (level 0 spells) can be cast unlimited times without consuming slots.

## Changes Made

### 1. Character Domain (internal/character/domain/character.go)
- Added `CurrentSpellSlots map[int]int` field to track remaining spell slots separately from max slots
- Updated `NewCharacter()` to initialize both `SpellSlots` (max) and `CurrentSpellSlots` (current available)
- Fixed case-insensitive class matching in `NewCharacter()`
- Added cantrip (Level 0) support for full caster classes
- Implemented `CastSpell()` method that:
  - Returns immediately for cantrips (level 0) without consuming slots
  - Checks if spell slots are available for the spell level
  - Consumes one spell slot when casting
  - Returns "No spell slot available!" error when no slots remain

### 2. Character Service (internal/character/service/character_service.go)
- Removed redundant `GetSpellSlots()` call (handled in `NewCharacter()` now)
- Added `CastSpell()` method that:
  - Loads the character
  - Validates the character can cast spells
  - Gets the spell level using `spell.GetSpellLevel()`
  - Calls character domain's `CastSpell()` method
  - Saves the updated character state

### 3. CLI Commands (internal/cli/commands.go)
- Implemented new `CastSpellCommand` with:
  - Command name: "cast-spell"
  - Required flags: -name and -spell
  - Displays updated spell slots after casting

### 4. Character Viewer (internal/cli/character_viewer.go)
- Updated spell slot display to show `CurrentSpellSlots` instead of max `SpellSlots`

### 5. Main Application (main.go)
- Registered the new `CastSpellCommand`

### 6. Web Template Data (internal/web/template_data.go)
- Added `CurrentSpellSlots` field to template data struct
- Populates current spell slots from character data

### 7. Spell Database (internal/spell/spell.go)
- Added "feeblemind" to the spell levels map as level 8

## Testing
Successfully tested the complete scenario from the assignment:
1. Created Gandalf (wizard, level 20)
2. Prepared fire bolt (cantrip), burning hands (level 1), and feeblemind (level 8)
3. Verified initial spell slots match expected output
4. Cast burning hands - Level 1 slots decreased from 4 to 3
5. Cast fire bolt - Cantrip didn't consume any slots
6. Cast feeblemind - Level 8 slot decreased from 1 to 0
7. Attempted to cast feeblemind again - Received "No spell slot available!" error

All outputs match the expected results from the assignment specification.

## Git Commit
All changes committed with message: "Implement spell slot usage tracking"
Git diff available in: exam_resit_spell_slots.patch
