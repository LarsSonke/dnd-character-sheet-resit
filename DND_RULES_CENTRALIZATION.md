# D&D Rules Centralization Analysis

## ✅ SINGLE SOURCE OF TRUTH FOR EACH RULE

### **Spell Slot Progression** (by class & level)
**Location**: `internal/character/domain/spell_slots.go`
- `FullCasterSpellSlots(level)` - Wizard, Cleric, Druid, Bard, Sorcerer
- `HalfCasterSpellSlots(level)` - Paladin, Ranger
- `PactMagicSpellSlots(level)` - Warlock
- `FullCasterCantrips(level)` - Cantrips for full casters

**Used by**: 
- `Character.GetSpellSlots()` ← Single point of use
- `NewCharacter()` calls `GetSpellSlots()` ← No duplication!

**Impact of change**: Modify spell_slots.go → Automatically applies everywhere

---

### **Spell Levels** (which spell is what level)
**Location**: `internal/spell/spell.go`
- `GetSpellLevel(spellName)` function
- Hardcoded map + CSV fallback

**Used by**:
- `CharacterService.PrepareSpell()` - validates spell level
- `CharacterService.CastSpell()` - determines which slot to consume
- `markdown_formatter.go` - formats spell lists

**Impact of change**: Add spell to map → Automatically available to all commands

---

### **Spell Slot Consumption** (cantrips don't consume, others do)
**Location**: `internal/character/domain/character.go`
- `Character.CastSpell(spellLevel)` method

**Used by**:
- `CharacterService.CastSpell()` ← Single caller

**Impact of change**: Modify CastSpell() → Behavior changes for all callers

---

### **Class Spellcasting Abilities**
**Location**: `internal/character/domain/character.go`
- `Character.GetSpellSlots()` - Maps classes to slot functions
- `Character.IsSpellcaster()` - Defines which classes can cast
- `Character.IsPreparedCaster()` - Defines prepared vs known casters
- `Character.SpellcastingAbility()` - INT/WIS/CHA per class

**Impact of change**: Single method change → All features updated

---

### **Ability Score Modifiers**
**Location**: `internal/character/domain/skills.go`
- `Modifier(score)` function

**Used by**: All ability checks, saving throws, spell attacks
**Impact of change**: Change formula → Applies everywhere

---

### **Proficiency Bonus**
**Location**: `internal/character/domain/character.go`
- `ProficiencyBonus(level)` function

**Impact of change**: Change formula → Applies to all proficient skills/saves

---

### **Racial Bonuses**
**Location**: `internal/character/domain/character.go`
- `Race.GetAbilityBonuses()` method

**Impact of change**: Add/modify race → Single location update

---

### **Background Skills**
**Location**: `internal/character/domain/character.go`
- `Background.GetSkillProficiencies()` method

**Impact of change**: Modify background → Single location update

---

### **Class Skills**
**Location**: `internal/character/domain/character.go`
- `Class.GetAvailableSkills()` method
- `Class.GetSkillCount()` method

**Impact of change**: Modify class skills → Single location update

---

## ✅ MAINTAINABILITY SCORE: EXCELLENT

### Open/Closed Principle Adherence:
✅ **Adding new spell**: Update `GetSpellLevel()` map only
✅ **Adding new class**: Update `GetSpellSlots()` switch only
✅ **Changing spell slot progression**: Update `spell_slots.go` only
✅ **Changing spell consumption logic**: Update `CastSpell()` only
✅ **Adding new race**: Update `GetAbilityBonuses()` only

### No Cascading Changes Required:
- ✅ All D&D rules in domain layer
- ✅ Single source of truth for each rule
- ✅ Service layer orchestrates, doesn't duplicate logic
- ✅ CLI layer presents, doesn't implement rules

### Architecture Compliance:
✅ Domain layer: Pure D&D business rules
✅ Service layer: Application orchestration
✅ CLI layer: User interface only
✅ Dependencies flow inward only

---

## Changes Made in This Refactoring:

**Before**: 
- Spell slot calculation duplicated in `NewCharacter()` AND `GetSpellSlots()`
- Potential for inconsistency if only one location updated

**After**:
- `NewCharacter()` calls `c.GetSpellSlots()` 
- Single source of truth: `GetSpellSlots()` method
- **Zero duplication of D&D rules**

**Impact**: If D&D spell slot rules change → Update one method → All code updated
