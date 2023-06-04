TAGS= forum thread user service post
OPENAPI=swagger.yml

.PHONY: $(TAGS)
$(TAGS):
	@mkdir -p internal/$@/delivery/http
	@mkdir -p internal/$@/repository/pgx
	@mkdir -p internal/$@/service
	@oapi-codegen -package=api -generate=types,fiber -include-tags=$@ $(OPENAPI) > internal/$@/delivery/http/handlers.gen.go

.PHONY: codegen
codegen: $(TAGS)
	@mkdir -p internal/models
	@oapi-codegen -package=models -generate=types,skip-prune -include-tags=$@ $(OPENAPI) > internal/models/models.gen.go