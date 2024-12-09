package day9

import (
	"goated-aoc-2024/year2024"
	"strconv"
	"strings"
)

// OptimizeContiguousFreeSpace algorithm:
// Move blocks one at a time from the end of the disk to the left most free spaceuntil there are no gaps
// Format of disk map is alternating single digits, first digit is number of blocks, followed by number of free
// space after that file. The ID number is the file natural ordering
func OptimizeContiguousFreeSpace(input string) int64 {
	originalFileMetadata := year2024.Deque[FileMetadata]{}
	split := strings.Split(input, "")
	id := 0
	currentFreeSpace := 0
	offset := 0
	for i := 0; i < len(split)-1; i += 2 {
		blockSize, _ := strconv.Atoi(split[i])
		freeSpace, _ := strconv.Atoi(split[i+1])
		if i == 0 {
			currentFreeSpace = freeSpace
			offset = blockSize
		} else {
			originalFileMetadata.Enqueue(FileMetadata{FileId: id, NumBlocks: blockSize, FollowingFreeSpace: freeSpace})
		}
		id++
	}
	lastBlockSize, _ := strconv.Atoi(split[len(split)-1])
	originalFileMetadata.Enqueue(FileMetadata{FileId: id, NumBlocks: lastBlockSize, FollowingFreeSpace: 0})
	fileMetadataToInsert, _ := originalFileMetadata.RemoveLast()
	checkSum := int64(0)
	for !originalFileMetadata.IsEmpty() {
		if currentFreeSpace == 0 {
			front, _ := originalFileMetadata.RemoveFirst()
			checkSum += evaluateCheckSum(front, offset)
			offset += front.NumBlocks
			currentFreeSpace = front.FollowingFreeSpace
		} else {
			if fileMetadataToInsert.NumBlocks == 0 {
				fileMetadataToInsert, _ = originalFileMetadata.RemoveLast()
			} else {
				blocksToInsert := min(fileMetadataToInsert.NumBlocks, currentFreeSpace)
				checkSum += int64(fileMetadataToInsert.FileId) * int64(sumRange(offset, offset+blocksToInsert-1))
				offset += blocksToInsert
				currentFreeSpace -= blocksToInsert
				fileMetadataToInsert.NumBlocks -= blocksToInsert
			}
		}
	}

	if fileMetadataToInsert.NumBlocks > 0 {
		checkSum += int64(fileMetadataToInsert.FileId) * int64(sumRange(offset, offset+fileMetadataToInsert.NumBlocks-1))
	}
	return checkSum
}

// OptimizeContiguousFreeSpaceWithoutFragmentation Key difference here -- only move file blocks if we can move all blocks of the file
// to avoid fragmented files
func OptimizeContiguousFreeSpaceWithoutFragmentation(input string) int64 {
	originalFileMetadata := year2024.Deque[FileMetadata]{}
	split := strings.Split(input, "")
	checkSum := int64(0)
	id := 0
	fullOffset := 0
	freeSpaceToOffset := make(map[int]*year2024.MinHeap)
	for i := 0; i < len(split)-1; i += 2 {
		blockSize, _ := strconv.Atoi(split[i])
		freeSpace, _ := strconv.Atoi(split[i+1])
		_, contains := freeSpaceToOffset[freeSpace]
		if !contains {
			freeSpaceToOffset[freeSpace] = &year2024.MinHeap{}
		}
		if i == 0 {
			fullOffset = blockSize + freeSpace
			freeSpaceToOffset[freeSpace].Offer(blockSize)
		} else {
			originalFileMetadata.Enqueue(FileMetadata{FileId: id, NumBlocks: blockSize, FollowingFreeSpace: freeSpace, Offset: fullOffset})
			freeSpaceToOffset[freeSpace].Offer(fullOffset + blockSize)
			fullOffset = fullOffset + blockSize + freeSpace
		}
		id++
	}
	lastBlockSize, _ := strconv.Atoi(split[len(split)-1])
	originalFileMetadata.Enqueue(FileMetadata{FileId: id, NumBlocks: lastBlockSize, FollowingFreeSpace: 0, Offset: fullOffset})
	for !originalFileMetadata.IsEmpty() {
		fileMetadataToInsert, _ := originalFileMetadata.RemoveLast()
		requiredFreeSpace := fileMetadataToInsert.NumBlocks
		anyMatch := false
		minimumOffset := fileMetadataToInsert.Offset
		freeSpaceOfMinimumOffset := 0
		for freeSpace := requiredFreeSpace; freeSpace <= 9; freeSpace++ {
			heap, contains := freeSpaceToOffset[freeSpace]
			if !contains || heap.IsEmpty() {
				continue
			}
			nextFreeSpaceOffset, _ := heap.Peek()
			if nextFreeSpaceOffset < minimumOffset {
				anyMatch = true
				minimumOffset = nextFreeSpaceOffset
				freeSpaceOfMinimumOffset = freeSpace
			}

		}
		if !anyMatch {
			checkSum += evaluateCheckSum(fileMetadataToInsert, fileMetadataToInsert.Offset)
		} else {
			heap, _ := freeSpaceToOffset[freeSpaceOfMinimumOffset]
			heap.Remove()
			checkSum += evaluateCheckSum(fileMetadataToInsert, minimumOffset)
			newOffset := minimumOffset + fileMetadataToInsert.NumBlocks
			remainingFreeSpace := freeSpaceOfMinimumOffset - requiredFreeSpace
			if remainingFreeSpace > 0 {
				freeSpaceToOffset[remainingFreeSpace].Offer(newOffset)
			}
		}
	}

	return checkSum
}

type FileMetadata struct {
	FileId             int
	NumBlocks          int
	FollowingFreeSpace int
	Offset             int
}

func sumRange(n, m int) int {
	return (m - n + 1) * (n + m) / 2
}

func evaluateCheckSum(fileMetadata FileMetadata, offset int) int64 {
	return int64(fileMetadata.FileId) * int64(sumRange(offset, offset+fileMetadata.NumBlocks-1))
}
