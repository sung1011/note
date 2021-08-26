# go archive

## code

- [tar](../script/go/package/archive_tar_test.go)
- [zip](../script/go/package/archive_zip_test.go)

## tar

`tar` 是一种打包格式, 不压缩, 所以打包速度快, 体积大

```go doc
package tar // import "archive/tar"

Package tar implements access to tar archives.

Tape archives (tar) are a file format for storing a sequence of files that
can be read and written in a streaming manner. This package aims to cover
most variations of the format, including those produced by GNU and BSD tar
tools.

CONSTANTS

// 实体文件类型
const (
    // Type '0' indicates a regular file.
    TypeReg  = '0'
    TypeRegA = '\x00' // Deprecated: Use TypeReg instead.

    // Type '1' to '6' are header-only flags and may not have a data body.
    TypeLink    = '1' // Hard link
    TypeSymlink = '2' // Symbolic link
    TypeChar    = '3' // Character device node
    TypeBlock   = '4' // Block device node
    TypeDir     = '5' // Directory
    TypeFifo    = '6' // FIFO node

    // Type '7' is reserved.
    TypeCont = '7'

    // Type 'x' is used by the PAX format to store key-value records that
    // are only relevant to the next file.
    // This package transparently handles these types.
    TypeXHeader = 'x'

    // Type 'g' is used by the PAX format to store key-value records that
    // are relevant to all subsequent files.
    // This package only supports parsing and composing such headers,
    // but does not currently support persisting the global state across files.
    TypeXGlobalHeader = 'g'

    // Type 's' indicates a sparse file in the GNU format.
    TypeGNUSparse = 's'

    // Types 'L' and 'K' are used by the GNU format for a meta file
    // used to store the path or link name for the next file.
    // This package transparently handles these types.
    TypeGNULongName = 'L'
    TypeGNULongLink = 'K'
)
    Type flags for Header.Typeflag.


VARIABLES

var (
    ErrHeader          = errors.New("archive/tar: invalid tar header")
    ErrWriteTooLong    = errors.New("archive/tar: write too long")
    ErrFieldTooLong    = errors.New("archive/tar: header field too long")
    ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

TYPES

type Format int
    Format represents the tar archive format.

    The original tar format was introduced in Unix V7. Since then, there have
    been multiple competing formats attempting to standardize or extend the V7
    format to overcome its limitations. The most common formats are the USTAR,
    PAX, and GNU formats, each with their own advantages and limitations.

    The following table captures the capabilities of each format:

                          |  USTAR |       PAX |       GNU
        ------------------+--------+-----------+----------
        Name              |   256B | unlimited | unlimited
        Linkname          |   100B | unlimited | unlimited
        Size              | uint33 | unlimited |    uint89
        Mode              | uint21 |    uint21 |    uint57
        Uid/Gid           | uint21 | unlimited |    uint57
        Uname/Gname       |    32B | unlimited |       32B
        ModTime           | uint33 | unlimited |     int89
        AccessTime        |    n/a | unlimited |     int89
        ChangeTime        |    n/a | unlimited |     int89
        Devmajor/Devminor | uint21 |    uint21 |    uint57
        ------------------+--------+-----------+----------
        string encoding   |  ASCII |     UTF-8 |    binary
        sub-second times  |     no |       yes |        no
        sparse files      |     no |       yes |       yes

    The table''s upper portion shows the Header fields, where each format reports
    the maximum number of bytes allowed for each string field and the integer
    type used to store each numeric field (where timestamps are stored as the
    number of seconds since the Unix epoch).

    The table''s lower portion shows specialized features of each format, such as
    supported string encodings, support for sub-second timestamps, or support
    for sparse files.

    The Writer currently provides no support for sparse files.

const (

    // FormatUnknown indicates that the format is unknown.
    FormatUnknown Format

    // FormatUSTAR represents the USTAR header format defined in POSIX.1-1988.
    //
    // While this format is compatible with most tar readers,
    // the format has several limitations making it unsuitable for some usages.
    // Most notably, it cannot support sparse files, files larger than 8GiB,
    // filenames larger than 256 characters, and non-ASCII filenames.
    //
    // Reference:
    //  http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html#tag_20_92_13_06
    FormatUSTAR

    // FormatPAX represents the PAX header format defined in POSIX.1-2001.
    //
    // PAX extends USTAR by writing a special file with Typeflag TypeXHeader
    // preceding the original header. This file contains a set of key-value
    // records, which are used to overcome USTAR''s shortcomings, in addition to
    // providing the ability to have sub-second resolution for timestamps.
    //
    // Some newer formats add their own extensions to PAX by defining their
    // own keys and assigning certain semantic meaning to the associated values.
    // For example, sparse file support in PAX is implemented using keys
    // defined by the GNU manual (e.g., "GNU.sparse.map").
    //
    // Reference:
    //  http://pubs.opengroup.org/onlinepubs/009695399/utilities/pax.html
    FormatPAX

    // FormatGNU represents the GNU header format.
    //
    // The GNU header format is older than the USTAR and PAX standards and
    // is not compatible with them. The GNU format supports
    // arbitrary file sizes, filenames of arbitrary encoding and length,
    // sparse files, and other features.
    //
    // It is recommended that PAX be chosen over GNU unless the target
    // application can only parse GNU formatted archives.
    //
    // Reference:
    //  https://www.gnu.org/software/tar/manual/html_node/Standard.html
    FormatGNU
)
    Constants to identify various tar formats.

func (f Format) String() string

// 归档包中单个文件实体的Header
type Header struct {
    // Typeflag is the type of header entry.
    // The zero value is automatically promoted to either TypeReg or TypeDir
    // depending on the presence of a trailing slash in Name.
    Typeflag byte

    Name     string // Name of file entry
    Linkname string // Target name of link (valid for TypeLink or TypeSymlink)

    Size  int64  // Logical file size in bytes
    Mode  int64  // Permission and mode bits
    Uid   int    // User ID of owner
    Gid   int    // Group ID of owner
    Uname string // User name of owner
    Gname string // Group name of owner

    // If the Format is unspecified, then Writer.WriteHeader rounds ModTime
    // to the nearest second and ignores the AccessTime and ChangeTime fields.
    //
    // To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
    // To use sub-second resolution, specify the Format as PAX.
    ModTime    time.Time // Modification time
    AccessTime time.Time // Access time (requires either PAX or GNU support)
    ChangeTime time.Time // Change time (requires either PAX or GNU support)

    Devmajor int64 // Major device number (valid for TypeChar or TypeBlock); 字符设备或块设备的主设备号
    Devminor int64 // Minor device number (valid for TypeChar or TypeBlock); 字符设备或块设备的次设备号

    // Xattrs stores extended attributes as PAX records under the
    // "SCHILY.xattr." namespace.
    //
    // The following are semantically equivalent:
    //  h.Xattrs[key] = value
    //  h.PAXRecords["SCHILY.xattr."+key] = value
    //
    // When Writer.WriteHeader is called, the contents of Xattrs will take
    // precedence over those in PAXRecords.
    //
    // Deprecated: Use PAXRecords instead.
    Xattrs map[string]string

    // PAXRecords is a map of PAX extended header records.
    //
    // User-defined records should have keys of the following form:
    //  VENDOR.keyword
    // Where VENDOR is some namespace in all uppercase, and keyword may
    // not contain the '=' character (e.g., "GOLANG.pkg.version").
    // The key and value should be non-empty UTF-8 strings.
    //
    // When Writer.WriteHeader is called, PAX records derived from the
    // other fields in Header take precedence over PAXRecords.
    PAXRecords map[string]string

    // Format specifies the format of the tar header.
    //
    // This is set by Reader.Next as a best-effort guess at the format.
    // Since the Reader liberally reads some non-compliant files,
    // it is possible for this to be FormatUnknown.
    //
    // If the format is unspecified when Writer.WriteHeader is called,
    // then it uses the first format (in the order of USTAR, PAX, GNU)
    // capable of encoding this Header (see Format).
    Format Format
}
    A Header represents a single header in a tar archive. Some fields may not be
    populated.

    For forward compatibility, users that retrieve a Header from Reader.Next,
    mutate it in some ways, and then pass it back to Writer.WriteHeader should
    do so by creating a new Header and copying the fields that they are
    interested in preserving.

// `os.FileInfo` 转 `tar.Header`
func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
    FileInfoHeader creates a partially-populated Header from fi. If fi describes
    a symlink, FileInfoHeader records link as the link target. If fi describes a
    directory, a slash is appended to the name.

    Since os.FileInfo''s Name method only returns the base name of the file it
    describes, it may be necessary to modify Header.Name to provide the full
    path name of the file.

// `tar.Header` 转 `os.FileInfo`
func (h *Header) FileInfo() os.FileInfo
    FileInfo returns an os.FileInfo for the Header.

type Reader struct {
    // Has unexported fields.
}
    Reader provides sequential access to the contents of a tar archive.
    Reader.Next advances to the next file in the archive (including the first),
    and then Reader can be treated as an io.Reader to access the file''s data.

// 从r中创建reader
func NewReader(r io.Reader) *Reader
    NewReader creates a new Reader reading from r.

// 指向tar文件中的下一个文件实体, err为io.EOF是终止
func (tr *Reader) Next() (*Header, error)
    // Next advances to the next entry in the tar archive. The Header.Size
    // determines how many bytes can be read for the next file. Any remaining data
    // in the current file is automatically discarded.

    // io.EOF is returned at the end of the input.

// 读取tar中当前实体
func (tr *Reader) Read(b []byte) (int, error)
    // Read reads from the current file in the tar archive. It returns (0, io.EOF)
    // when it reaches the end of that file, until Next is called to advance to the
    // next file.

    // If the current file is sparse, then the regions marked as a hole are read
    // back as NUL-bytes.

    // Calling Read on special types like TypeLink, TypeSymlink, TypeChar,
    // TypeBlock, TypeDir, and TypeFifo returns (0, io.EOF) regardless of what the
    // Header.Size claims.

type Writer struct {
    // Has unexported fields.
}
    Writer provides sequential writing of a tar archive. Write.WriteHeader
    begins a new file with the provided Header, and then Writer can be treated
    as an io.Writer to supply that file''s data.

func NewWriter(w io.Writer) *Writer
    NewWriter creates a new Writer writing to w.

func (tw *Writer) Close() error
    Close closes the tar archive by flushing the padding, and writing the
    footer. If the current file (from a prior call to WriteHeader) is not fully
    written, then this returns an error.

func (tw *Writer) Flush() error
    Flush finishes writing the current file''s block padding. The current file
    must be fully written before Flush can be called.

    This is unnecessary as the next call to WriteHeader or Close will implicitly
    flush out the file''s padding.

func (tw *Writer) Write(b []byte) (int, error)
    Write writes to the current file in the tar archive. Write returns the error
    ErrWriteTooLong if more than Header.Size bytes are written after
    WriteHeader.

    Calling Write on special types like TypeLink, TypeSymlink, TypeChar,
    TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless of
    what the Header.Size claims.

func (tw *Writer) WriteHeader(hdr *Header) error
    WriteHeader writes hdr and prepares to accept the file''s contents. The
    Header.Size determines how many bytes can be written for the next file. If
    the current file is not fully written, then this returns an error. This
    implicitly flushes any padding necessary before writing the header.
```

## zip

```go
package zip // import "archive/zip"

Package zip provides support for reading and writing ZIP archives.

See: https://www.pkware.com/appnote

This package does not support disk spanning.

A note about ZIP64:

To be backwards compatible the FileHeader has both 32 and 64 bit Size
fields. The 64 bit fields will always contain the correct value and for
normal archives both fields will be the same. For files requiring the ZIP64
format the 32 bit fields will be 0xffffffff and the 64 bit fields must be
used instead.

CONSTANTS

const (
    Store   uint16 = 0 // no compression
    Deflate uint16 = 8 // DEFLATE compressed
)
    Compression methods.


VARIABLES

var (
    ErrFormat    = errors.New("zip: not a valid zip file")
    ErrAlgorithm = errors.New("zip: unsupported compression algorithm")
    ErrChecksum  = errors.New("zip: checksum error")
)

FUNCTIONS

func RegisterCompressor(method uint16, comp Compressor)
    RegisterCompressor registers custom compressors for a specified method ID.
    The common methods Store and Deflate are built in.

func RegisterDecompressor(method uint16, dcomp Decompressor)
    RegisterDecompressor allows custom decompressors for a specified method ID.
    The common methods Store and Deflate are built in.


TYPES

type Compressor func(w io.Writer) (io.WriteCloser, error)
    A Compressor returns a new compressing writer, writing to w. The
    WriteCloser''s Close method must be used to flush pending data to w. The
    Compressor itself must be safe to invoke from multiple goroutines
    simultaneously, but each returned writer will be used only by one goroutine
    at a time.

type Decompressor func(r io.Reader) io.ReadCloser
    A Decompressor returns a new decompressing reader, reading from r. The
    ReadCloser''s Close method must be used to release associated resources. The
    Decompressor itself must be safe to invoke from multiple goroutines
    simultaneously, but each returned reader will be used only by one goroutine
    at a time.

type File struct {
    FileHeader

    // Has unexported fields.
}

func (f *File) DataOffset() (offset int64, err error)
    DataOffset returns the offset of the file''s possibly-compressed data,
    relative to the beginning of the zip file.

    Most callers should instead use Open, which transparently decompresses data
    and verifies checksums.

func (f *File) Open() (io.ReadCloser, error)
    Open returns a ReadCloser that provides access to the File''s contents.
    Multiple files may be read concurrently.

type FileHeader struct {
    // Name is the name of the file.
    //
    // It must be a relative path, not start with a drive letter (such as "C:"),
    // and must use forward slashes instead of back slashes. A trailing slash
    // indicates that this file is a directory and should have no data.
    //
    // When reading zip files, the Name field is populated from
    // the zip file directly and is not validated for correctness.
    // It is the caller''s responsibility to sanitize it as
    // appropriate, including canonicalizing slash directions,
    // validating that paths are relative, and preventing path
    // traversal through filenames ("../../../").
    Name string

    // Comment is any arbitrary user-defined string shorter than 64KiB.
    Comment string

    // NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
    //
    // By specification, the only other encoding permitted should be CP-437,
    // but historically many ZIP readers interpret Name and Comment as whatever
    // the system''s local character encoding happens to be.
    //
    // This flag should only be set if the user intends to encode a non-portable
    // ZIP file for a specific localized region. Otherwise, the Writer
    // automatically sets the ZIP format''s UTF-8 flag for valid UTF-8 strings.
    NonUTF8 bool

    CreatorVersion uint16
    ReaderVersion  uint16
    Flags          uint16

    // Method is the compression method. If zero, Store is used.
    Method uint16

    // Modified is the modified time of the file.
    //
    // When reading, an extended timestamp is preferred over the legacy MS-DOS
    // date field, and the offset between the times is used as the timezone.
    // If only the MS-DOS date is present, the timezone is assumed to be UTC.
    //
    // When writing, an extended timestamp (which is timezone-agnostic) is
    // always emitted. The legacy MS-DOS date field is encoded according to the
    // location of the Modified time.
    Modified     time.Time
    ModifiedTime uint16 // Deprecated: Legacy MS-DOS date; use Modified instead.
    ModifiedDate uint16 // Deprecated: Legacy MS-DOS time; use Modified instead.

    CRC32              uint32
    CompressedSize     uint32 // Deprecated: Use CompressedSize64 instead.
    UncompressedSize   uint32 // Deprecated: Use UncompressedSize64 instead.
    CompressedSize64   uint64
    UncompressedSize64 uint64
    Extra              []byte
    ExternalAttrs      uint32 // Meaning depends on CreatorVersion
}
    FileHeader describes a file within a zip file. See the zip spec for details.

func FileInfoHeader(fi os.FileInfo) (*FileHeader, error)
    FileInfoHeader creates a partially-populated FileHeader from an os.FileInfo.
    Because os.FileInfo''s Name method returns only the base name of the file it
    describes, it may be necessary to modify the Name field of the returned
    header to provide the full path name of the file. If compression is desired,
    callers should set the FileHeader.Method field; it is unset by default.

func (h *FileHeader) FileInfo() os.FileInfo
    FileInfo returns an os.FileInfo for the FileHeader.

func (h *FileHeader) ModTime() time.Time
    ModTime returns the modification time in UTC using the legacy ModifiedDate
    and ModifiedTime fields.

    Deprecated: Use Modified instead.

func (h *FileHeader) Mode() (mode os.FileMode)
    Mode returns the permission and mode bits for the FileHeader.

func (h *FileHeader) SetModTime(t time.Time)
    SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields to the
    given time in UTC.

    Deprecated: Use Modified instead.

func (h *FileHeader) SetMode(mode os.FileMode)
    SetMode changes the permission and mode bits for the FileHeader.

type ReadCloser struct {
    Reader
    // Has unexported fields.
}

func OpenReader(name string) (*ReadCloser, error)
    OpenReader will open the Zip file specified by name and return a ReadCloser.

func (rc *ReadCloser) Close() error
    Close closes the Zip file, rendering it unusable for I/O.

type Reader struct {
    File    []*File
    Comment string
    // Has unexported fields.
}

func NewReader(r io.ReaderAt, size int64) (*Reader, error)
    NewReader returns a new Reader reading from r, which is assumed to have the
    given size in bytes.

func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
    RegisterDecompressor registers or overrides a custom decompressor for a
    specific method ID. If a decompressor for a given method is not found,
    Reader will default to looking up the decompressor at the package level.

type Writer struct {
    // Has unexported fields.
}
    Writer implements a zip file writer.

func NewWriter(w io.Writer) *Writer
    NewWriter returns a new Writer writing a zip file to w.

func (w *Writer) Close() error
    Close finishes writing the zip file by writing the central directory. It
    does not close the underlying writer.

func (w *Writer) Create(name string) (io.Writer, error)
    Create adds a file to the zip file using the provided name. It returns a
    Writer to which the file contents should be written. The file contents will
    be compressed using the Deflate method. The name must be a relative path: it
    must not start with a drive letter (e.g. C:) or leading slash, and only
    forward slashes are allowed. To create a directory instead of a file, add a
    trailing slash to the name. The file''s contents must be written to the
    io.Writer before the next call to Create, CreateHeader, or Close.

func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
    CreateHeader adds a file to the zip archive using the provided FileHeader
    for the file metadata. Writer takes ownership of fh and may mutate its
    fields. The caller must not modify fh after calling CreateHeader.

    This returns a Writer to which the file contents should be written. The
    file''s contents must be written to the io.Writer before the next call to
    Create, CreateHeader, or Close.

func (w *Writer) Flush() error
    Flush flushes any buffered data to the underlying writer. Calling Flush is
    not normally necessary; calling Close is sufficient.

func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
    RegisterCompressor registers or overrides a custom compressor for a specific
    method ID. If a compressor for a given method is not found, Writer will
    default to looking up the compressor at the package level.

func (w *Writer) SetComment(comment string) error
    SetComment sets the end-of-central-directory comment field. It can only be
    called before Close.

func (w *Writer) SetOffset(n int64)
    SetOffset sets the offset of the beginning of the zip data within the
    underlying writer. It should be used when the zip data is appended to an
    existing file, such as a binary executable. It must be called before any
    data is written.

```