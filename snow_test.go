package snowflake

import "testing"

func BenchmarkNewSnowflake(b *testing.B) {
    a, _ := NewSnowflake(0, 0)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := a.nextId()
        if err != nil {
            b.Error(err)
        }
    }
}

func BenchmarkNewSnowflake_Lock(b *testing.B) {
    a, _ := NewSnowflake(0, 0)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := a.NextId()
        if err != nil {
            b.Error(err)
        }
    }
}
