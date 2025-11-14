package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MdmBrands holds the schema definition for the MdmBrands entity.
type MdmBrands struct {
	ent.Schema
}

// Fields of the MdmBrands.
func (MdmBrands) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("brand_id"), // 将Ent的id字段映射到数据库的brand_id列

		field.String("brand_name").
			MaxLen(50).Comment("品牌名称").
			NotEmpty(),

		field.String("brand_logo").
			MaxLen(255).Comment("品牌logo222").Default(""),

		field.String("brand_title").
			MaxLen(255).Comment("品牌标题").Default(""),
	}
}

// Edges of the MdmBrands.
func (MdmBrands) Edges() []ent.Edge {
	return nil
}
func (MdmBrands) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment("基础物料品牌222"),
	}
}
