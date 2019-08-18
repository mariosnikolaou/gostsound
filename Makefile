SUBDIRS := StSoundLibrary StSoundLibrary/lzh
SUBCLEAN = $(addsuffix .clean,$(SUBDIRS))

.PHONY: all clean $(SUBDIRS) $(SUBCLEAN)

all: $(SUBDIRS)
		go build

clean: $(SUBCLEAN)

$(SUBDIRS):
		$(MAKE) -C $@

$(SUBCLEAN): %.clean:
		$(MAKE) -C $* clean
		rm -f gostsound
