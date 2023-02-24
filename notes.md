A CLI for facilitating image/text smilarity search using pgvector

neon-vectorize
Subcommands
 - init: creates db schema in a Neon project, creates extension pgvector
 - seed: takes directory of files and inserts them into data table
 - vectorize: processes datatable to make embeddings and insert them to embedding table
   args: should accept type of convertion, could use OpenAI or replicate.com or pg extension
 - find: tries to find closest image/text and return it
   args: input file
