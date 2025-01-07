package stac

// Constant to define the STAC specification version
const STAC_VERSION = "1.1.0"

// Constant to define the STAC relationship type for 'self'
const STAC_LINK_RELTYPE_SELF = "self"

// Constant to define the STAC relationship type for 'root'
const STAC_LINK_RELTYPE_ROOT = "root"

// Constant to define the STAC relationship type for 'parent'
const STAC_LINK_RELTYPE_PARENT = "parent"

// Constant to define the STAC relationship type for 'child'
const STAC_LINK_RELTYPE_CHILD = "child"

// Constant to define the STAC relationship type for 'collection'
const STAC_LINK_RELTYPE_COLLECTION = "collection"

// Constant to define the STAC relationship type for 'item'
const STAC_LINK_RELTYPE_ITEM = "item"

// Constant to define the schema location for the STAC Table extension
const STAC_EXTENSIONS_TABLE = "https://stac-extensions.github.io/table/v1.2.0/schema.json"

// Constant to define the STAC 'type' for "Catalog"
const STAC_TYPES_CATALOG = "Catalog"

// Constant to define the STAC 'type' for "Collection"
const STAC_TYPES_COLLECTION = "Collection"

// Constant to define the STAC 'type' for "Item"
const STAC_TYPES_ITEM = "Feature"

// Constant to define the STAC Asset 'role' for "Thumbnail"
const STAC_ASSET_ROLE_THUMBNAIL = "thumbnail"

// Constant to define the STAC Asset 'role' for "Overview"
const STAC_ASSET_ROLE_OVERVIEW = "overview"

// Constant to define the STAC Asset 'role' for "Data"
const STAC_ASSET_ROLE_DATA = "data"

// Constant to define the STAC Asset 'role' for "Metadata"
const STAC_ASSET_ROLE_METADATA = "metadata"

// Constant to define the IANA Media Type for GeoTIFF
const STAC_ASSET_MEDIA_TYPE_GEOTIFF = "image/tiff; application=geotiff"

// Constant to define the IANA Media Type for Cloud Optimized GeoTIFF
const STAC_ASSET_MEDIA_TYPE_CLOUDOPTIMIZEDGEOTIFF = "image/tiff; application=geotiff; profile=cloud-optimized"

// Constant to define the IANA Media Type for JPEG2000
const STAC_ASSET_MEDIA_TYPE_JPEG2000 = "image/jp2"

// Constant to define the IANA Media Type for PNG
const STAC_ASSET_MEDIA_TYPE_PNG = "image/png"

// Constant to define the IANA Media Type for JPEG
const STAC_ASSET_MEDIA_TYPE_JPEG = "image/jpeg"

// Constant to define the IANA Media Type for XML
const STAC_ASSET_MEDIA_TYPE_XML = "application/xml"

// Constant to define the IANA Media Type for JSON
const STAC_ASSET_MEDIA_TYPE_JSON = "application/json"

// Constant to define the IANA Media Type for Plain Text
const STAC_ASSET_MEDIA_TYPE_PLAINTEXT = "text/plain"

// Constant to define the IANA Media Type for GeoJSON
const STAC_ASSET_MEDIA_TYPE_GEOJSON = "application/geo+json"

// Constant to define the IANA Media Type for GeoPackage
const STAC_ASSET_MEDIA_TYPE_GEOPACKAGE = "application/geopackage+sqlite3"

// Constant to define the IANA Media Type for HDF5
const STAC_ASSET_MEDIA_TYPE_HDF5 = "application/x-hdf5"

// Constant to define the IANA Media Type for HDF
const STAC_ASSET_MEDIA_TYPE_HDF4 = "application/x-hdf"

// Constant to define the IANA Media Type for Cloud OptimizedPoint Cloud
const STAC_ASSET_MEDIA_TYPE_CLOUDOPTIMIZEDPOINTCLOUD = "application/vnd.laszip+copc"

// Constant to define the IANA Media Type for Apache GeoParquet
const STAC_ASSET_MEDIA_TYPE_APACHEGEOPARQUET = "application/vnd.apache.parquet"

// Constant to define the IANA Media Type for OGC 3D Tiles
const STAC_ASSET_MEDIA_TYPE_OGC3DTILES = "application/3dtiles+json"

// Constant to define the IANA Media Type for Protomaps PMTiles
const STAC_ASSET_MEDIA_TYPE_PROTOMAPSPMTILES = "application/vnd.pmtiles"

// Constant to define the IANA Media Type for NetCDF
const STAC_ASSET_MEDIA_TYPE_NETCDF = "application/netcdf"

// Constant to define the GeoJSON feature type for 'Point'
const GEOJSON_TYPE_POINT = "Point"

// Constant to define the GeoJSON feature type for 'LineString'
const GEOJSON_TYPE_LINESTRING = "LineString"

// Constant to define the GeoJSON feature type for 'Polygon'
const GEOJSON_TYPE_POLYGON = "Polygon"

// Contant to define the IANA Media Type for STAC Items
const STAC_ITEM_MIME_TYPE = STAC_ASSET_MEDIA_TYPE_GEOJSON

// Contant to define the IANA Media Type for STAC Catalogs
const STAC_CATALOG_MIME_TYPE = STAC_ASSET_MEDIA_TYPE_JSON

// Contant to define the IANA Media Type for STAC Collections
const STAC_COLLECTION_MIME_TYPE = STAC_ASSET_MEDIA_TYPE_JSON
