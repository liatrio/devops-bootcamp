# Default to in-memory repo. To switch backend, change this import only.
from .in_memory_repo import get_all, add

# To use sqlite, replace above with:
# from .sqlite_repo import get_all, add
