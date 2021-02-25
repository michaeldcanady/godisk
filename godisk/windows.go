package godisk

import (
	"fmt"

	slicetils "github.com/michaeldcanady/Slicetils/SliceTils"
	"github.com/michaeldcanady/Windows-api/Windows-Api/kernel32"
	convert "github.com/michaeldcanady/goconversion/Conversion"
)

type Disk disk

type disk struct {
	volume                   string
	label                    string
	driveType                string
	nVolumeNameSize          uint32
	SerialNumber             uint32
	lpMaximumComponentLength uint32
	SystemFlags              []string
	FileSystem               string
	nFileSystemNameSize      uint32
	Free                     int64
	Used                     int64
	Total                    int64
}

func New() *Disk {
	D := Disk{}

	return &D
}

func GetDrive(volumeName string) (disk Disk, err error) {
	drives := kernel32.GetLogicalDriveStrings()

	if contains, _ := slicetils.Contains(drives, volumeName); !contains {
		return Disk{}, fmt.Errorf("%v is not a valid volume", volumeName)
	}

	typ, err := kernel32.GetDriveTypeW(volumeName)
	if err != nil {
		fmt.Println(err)
	}
	volumeInfo := kernel32.GetVolumeInformationW(volumeName)

	free, total, _, used := kernel32.GetDiskFreeSpaceEx(volumeName)
	return newDisk(volumeName, typ, volumeInfo, free, used, total), err
}

func newDisk(volume, driveType string, volumeInfo kernel32.Volume, Free, Used, Total int64) Disk {
	d := disk{
		volume:                   volume,
		label:                    volumeInfo.VolumeLabel,
		driveType:                driveType,
		SerialNumber:             volumeInfo.SerialNumber,
		lpMaximumComponentLength: volumeInfo.LpMaximumComponentLength,
		SystemFlags:              volumeInfo.SystemFlags,
		FileSystem:               volumeInfo.FileSystem,
		Free:                     Free,
		Used:                     Used,
		Total:                    Total,
	}
	return Disk(d)
}

func GetDrives() (disks []Disk) {
	drives := kernel32.GetLogicalDriveStrings()

	for _, d := range drives {
		typ, err := kernel32.GetDriveTypeW(d)
		if err != nil {
			fmt.Println(err)
		}
		volumeInfo := kernel32.GetVolumeInformationW(d)

		free, total, _, used := kernel32.GetDisckFreeSpaceEx(d)

		disks = append(disks, newDisk(d, typ, volumeInfo, free, used, total))
	}
	return disks
}

func (D Disk) String() string {
	return fmt.Sprintf("Volume:                  %v\n"+
		"Drive Type:              %v\n"+
		"VolumeLabel:             %v\n"+
		"Serial Number            %v\n"+
		"lpMaximumComponentLength %v\n"+
		"System Flags             %v\n"+
		"File System              %v\n"+
		"Free Space:              %v\n"+
		"Used Space:              %v\n"+
		"Total Space:             %v",
		D.volume, D.label, D.driveType,
		D.SerialNumber,
		D.lpMaximumComponentLength,
		D.SystemFlags,
		D.FileSystem,
		convert.ByteCountSI(D.Free, 0), convert.ByteCountSI(D.Used, 0), convert.ByteCountSI(D.Total, 0))
}
