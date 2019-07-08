package snowflake

import "testing"

func BenchmarkNewSnowflake(b *testing.B) {
    a, _ := NewSnowflake(0, 0)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := a.nextID()
        if err != nil {
            b.Error(err)
        }
    }
}

func BenchmarkNewSnowflake_Lock(b *testing.B) {
    a, _ := NewSnowflake(0, 0)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := a.NextID()
        if err != nil {
            b.Error(err)
        }
    }
}
