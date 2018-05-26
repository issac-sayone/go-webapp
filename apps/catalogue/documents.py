from oscar.apps.catalogue.models import Product

from django_elasticsearch_dsl import DocType, Index


product = Index('product')
# See Elasticsearch Indices API reference for available settings

product.settings(
    number_of_shards=1,
    number_of_replicas=0
)

@product.doc_type
class ProductDocument(DocType):

    class Meta:
        model = Product
        fields = [
            'title',
            'id'
        ]