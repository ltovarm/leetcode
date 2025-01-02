# Specification of file paths and names
SRCDIR := .
BUILDDIR := build
TARGET := $(BUILDDIR)/Length_of_last_word
SRCEXT := cpp

# Compiler options
CXX := g++
CXXFLAGS := -std=c++11 -Wall -Wextra -pedantic

# List of source files
SOURCES := $(wildcard $(SRCDIR)/*.$(SRCEXT))

# Objects generated from source files
OBJECTS := $(patsubst $(SRCDIR)/%,$(BUILDDIR)/%,$(SOURCES:.$(SRCEXT)=.o))

# Command to create the build directory if it doesn't exist
MKDIR_P = mkdir -p

# Rule for building the binary
$(TARGET): $(OBJECTS)
	@$(MKDIR_P) $(@D)
	$(CXX) $^ -o $(TARGET)

# Rule for building the objects
$(BUILDDIR)/%.o: $(SRCDIR)/%.$(SRCEXT)
	@$(MKDIR_P) $(@D)
	$(CXX) $(CXXFLAGS) -c -o $@ $<

# Cleaning
clean:
	$(RM) -r $(BUILDDIR)

# "Phony target" rule to avoid confusion with files named "all" or "clean"
.PHONY: all clean

# By default, build the binary
all: $(TARGET)
