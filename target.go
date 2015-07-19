// This package implemets a simulated network structure.
package FakeNet

// Target is an interface that will enforce types to be able to recieve packets.
type Target interface {
    recieve(p *Packet) error
}
